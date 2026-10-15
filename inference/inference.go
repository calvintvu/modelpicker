package inference

import (
	"fmt"
	"log"
	"math"
	"slices"

	"github.com/daulet/tokenizers"
	ort "github.com/yalue/onnxruntime_go"
)

// tokenizer.json
const (
	clsID  = 50281
	sepID  = 50282
	maskID = 50284
)

// laya_config.json max_len
const maxLen = 512

type InferenceSession struct {
	tokenizer *tokenizers.Tokenizer
	session   *ort.DynamicAdvancedSession
}

func LoadInferenceSession() (*InferenceSession, error) {
	ort.SetSharedLibraryPath("/usr/lib/libonnxruntime.so")
	log.Println("onnx runtime path set...")
	err := ort.InitializeEnvironment()
	if err != nil {
		return nil, fmt.Errorf("Error initializing onnxruntime library: %w", err)
	}
	log.Println("onnx runtime environment set...")
	tokenizer, err := tokenizers.FromFile("./inference/model/tokenizer/tokenizer.json")
	if err != nil {
		ort.DestroyEnvironment()
		return nil, fmt.Errorf("loading tokenizer: %w", err)
	}
	log.Println("tokenizer set...")
	session, err := ort.NewDynamicAdvancedSession("inference/model/laya.onnx",
		[]string{"input_ids", "attention_mask", "marker_pos", "marker_mask", "qtype"},
		[]string{"logits", "act_probs"}, nil)
	if err != nil {
		tokenizer.Close()
		ort.DestroyEnvironment()
		return nil, fmt.Errorf("loading model: %w", err)
	}
	log.Println("inference session loaded successfully!...")
	return &InferenceSession{tokenizer: tokenizer, session: session}, nil
}

func (l *InferenceSession) CloseInferenceSession() {
	l.session.Destroy()
	l.tokenizer.Close()
	ort.DestroyEnvironment()
}

func (session *InferenceSession) head(instructions string, options []string) (seq, markers []int64) {
	seq = append([]int64{clsID}, session.encode("choice question: "+instructions)...)
	seq = append(seq, sepID)
	for _, o := range options {
		markers = append(markers, int64(len(seq)))
		seq = append(append(seq, maskID), session.encode(" "+o)...)
	}
	return seq, markers
}

func (session *InferenceSession) CheckFits(instructions string, options []string) error {
	seq, _ := session.head(instructions, options)
	if len(seq)+2 >= maxLen {
		return fmt.Errorf("question + options use %d tokens, over max_len %d", len(seq), maxLen)
	}
	return nil
}

func (session *InferenceSession) Choose(instructions, state string, options []string, temperature float64) ([]float64, error) {
	seq, markers := session.head(instructions, options)
	room := maxLen - len(seq) - 2
	if room <= 0 {
		return nil, fmt.Errorf("question + options use %d tokens, over max_len %d", len(seq), maxLen)
	}
	stateIDs := session.encode(state)
	if len(stateIDs) > room {
		stateIDs = stateIDs[:room]
	}
	seq = append(append(append(seq, sepID), stateIDs...), sepID)

	n, k := int64(len(seq)), int64(len(options))
	inputs := []ort.Value{
		newTensor(seq, 1, n),
		newTensor(slices.Repeat([]int64{1}, int(n)), 1, n),
		newTensor(markers, 1, k),
		newTensor(slices.Repeat([]bool{true}, int(k)), 1, k),
		newTensor([]int64{0}, 1), // qtype 0 = choice
	}
	outputs := []ort.Value{nil, nil}
	defer func() {
		for _, v := range slices.Concat(inputs, outputs) {
			if v != nil {
				v.Destroy()
			}
		}
	}()
	if err := session.session.Run(inputs, outputs); err != nil {
		return nil, fmt.Errorf("running model: %w", err)
	}
	return softmax(outputs[0].(*ort.Tensor[float32]).GetData(), temperature), nil
}

func (session *InferenceSession) encode(s string) []int64 {
	ids, _ := session.tokenizer.Encode(s, false)
	out := make([]int64, len(ids))
	for i, id := range ids {
		out[i] = int64(id)
	}
	return out
}

func newTensor[T ort.TensorData](data []T, dims ...int64) ort.Value {
	t, err := ort.NewTensor(ort.NewShape(dims...), data)
	if err != nil {
		panic(err)
	}
	return t
}

func softmax(scores []float32, temperature float64) []float64 {
	out := make([]float64, len(scores))
	maxScore := float64(slices.Max(scores)) / temperature
	var sum float64
	for i, score := range scores {
		out[i] = math.Exp(float64(score)/temperature - maxScore)
		sum += out[i]
	}
	for i := range out {
		out[i] /= sum
	}
	return out
}
