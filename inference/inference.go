package inference

import (
	"fmt"
	"log"

	"github.com/daulet/tokenizers"
	ort "github.com/yalue/onnxruntime_go"
)

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
	log.Println("inference session set...")
	return &InferenceSession{tokenizer: tokenizer, session: session}, nil
}

func (l *InferenceSession) CloseInferenceSession() {
	l.session.Destroy()
	l.tokenizer.Close()
	ort.DestroyEnvironment()
}
