# Selection evaluations

These suites test model selection through `POST /v1/pick`. They do not test generated responses or provider quality.

## Run the baseline

Requires Python 3.11+ and the [native build tools](../README.md#native-setup).
Python dependencies are not required. The `hf` CLI is necessary only when the model files are missing.

From the repository root, run:

```sh
python3 evals/run.py
```

The runner installs missing native libraries through `make deps` and downloads missing weights through `make model`.
It reuses the existing model cache by default.
It builds a temporary binary, starts a private server, waits for readiness, and sends one authenticated warmup request.
Then it evaluates each case sequentially and stops the server.
Temporary build files are removed after success, failure, Ctrl+C, or SIGTERM.

The server uses a private working directory, a fresh API key, and an available loopback port.
The repository's `.env` and `config.yaml` do not control this server. Existing servers are not stopped.
Cached model files and native libraries remain available for later runs.

The [eval config](config.json) defines `tier_1`, `tier_2`, and `tier_3` independently of the normal service policy.
Its model IDs are placeholders for membership checks. The service never calls model providers.
Inference fallback is disabled so failures remain visible.
Edit this JSON file to change the classifier question or tier descriptions before a new evaluation.
The service's existing YAML loader also accepts JSON.

Optional paths and timeouts:

```sh
python3 evals/run.py --cases evals/suites/mt-bench-baseline.csv \
  --config evals/config.json --model-dir inference/model \
  --output /tmp/modelpicker-eval-result \
  --startup-timeout 180 --request-timeout 60
```

The output directory must not exist. The default creates a unique directory under `evals/runs/`.
Startup timeout starts after the server launches. Build and download time are separate.

Each run retains:

| File | Contents |
| --- | --- |
| `cases.csv`, `config.json` | Exact input snapshots. |
| `build.log`, `server.log` | Setup, build, startup, request, and shutdown logs. |
| `results.jsonl` | Each request, expected tier, predicted tier/model, latency, HTTP status, correctness, and errors. |
| `summary.json` | Run status, progress, input/model/runtime hashes, server exit status, and metrics. |

Metrics include accuracy, per-tier precision/recall/F1, macro F1, a confusion matrix, suite scores, and HTTP latency p50/p95.
The confusion matrix contains tier predictions. `errors_by_tier` records request failures separately.
Request failures and invalid responses count as incorrect. Warmup is excluded from the metrics.
Incomplete runs retain partial results and omit aggregate metrics.
The API key is not saved in the artifacts.

Exit codes: `0` means the run completed without request or lifecycle errors, `2` means failure, and `130` means interruption.
Tier mismatches do not cause a failed exit because no accuracy threshold is agreed.

The fixed baseline fits the eval config: all 100 prompts use at most 371 of its 414 available prompt tokens.
Custom cases or configurations need a fresh tokenizer audit because the service silently truncates input beyond its token budget.

## Suites

| Suite | Source | Focus | Status |
| --- | --- | --- | --- |
| MT-Bench baseline | [MT-Bench](https://github.com/lm-sys/FastChat/tree/main/fastchat/llm_judge/data/mt_bench) | Minimum required capability, independent of task category. | 80 first-turn prompts manually labeled. |
| Advanced supplement | [ResearchRubrics](https://huggingface.co/datasets/ScaleAI/researchrubrics) and [FrontierCS](https://huggingface.co/datasets/FrontierCS/Frontier-CS) | AI research and systems engineering. | 20 prompts appended to the baseline CSV; manually labeled. |
| General tasks | [WildBench](https://huggingface.co/datasets/allenai/WildBench) | Code, writing, reasoning, and general task mappings. | Planned. |
| Small talk | [CLINC150](https://huggingface.co/datasets/clinc/clinc_oos) | Greeting, goodbye, and thank-you test examples mapped to nano. | Planned. |
| Creative writing | [WritingBench](https://github.com/X-PLUG/WritingBench) | 42 English creative-writing requests mapped to writing. | Cases prepared; not run. |

## MT-Bench baseline labeling

The [CSV](suites/mt-bench-baseline.csv) contains 100 prompts: the original 80 MT-Bench prompts followed by 20 advanced requests.
The first 80 rows retain all user labels and notes.
All 100 rows have valid user labels: 27 `tier_1`, 51 `tier_2`, and 22 `tier_3`.
MT-Bench question IDs, categories, and prompt text remain unchanged, including line breaks and quotation marks.
The second turns are excluded because they depend on previous answers.
This file supports a derived routing evaluation, not an official MT-Bench generation score.

Fill `expected_tier` with one of these labels:

| Label | Capability | Minimum required work |
| --- | --- | --- |
| `tier_1` | Basic | Direct extraction, simple rewriting, straightforward questions, or routine tasks with few constraints. |
| `tier_2` | Standard | Normal drafting, coding, explanation, or analysis with moderate reasoning and several constraints. |
| `tier_3` | Advanced | Difficult reasoning, complex debugging or algorithms, demanding synthesis, or interacting constraints that require careful reasoning. |

Choose the lowest tier you expect can give a correct, relevant answer that follows the prompt's instructions.
Use the same quality standard across categories.
Do not assign a tier from the topic or prompt length alone.
If you cannot decide, leave `expected_tier` blank and explain the uncertainty in `notes`.
The `category` column records the upstream task category, not the expected tier.

These labels describe required capability. Model cost bands can map to them later, after the model pools are defined.
Parameter count and price alone do not establish capability.
Manual labels record the intended routing policy. Model answer evaluations must establish whether each pool meets that policy.
MT-Bench contains challenging requests, so this baseline does not guarantee balanced capability tiers or enough basic tasks.

The normal `config.yaml` uses task categories as tiers. The runner uses the matching capability tiers in `evals/config.json` instead.
The labeling file does not change the service configuration or call the picker.

Source revision: `b494d0c6b4e7935f1764f8439e75da3e66beccc7`.
Source file: `fastchat/llm_judge/data/mt_bench/question.jsonl`.
Source SHA-256: `119565adbab82227089cefdb44c8d7e2cf04dc0a0ec233634c82e7d4e2a944f7`.

## Advanced supplement

The final 20 CSV rows contain 10 AI research requests and 10 systems engineering requests.
Four full prompts come from ResearchRubrics and 16 task excerpts come from FrontierCS.
The [source manifest](suites/advanced-source.json) records pinned revisions, source hashes, original IDs, topics, and category mappings.
The excerpts retain source wording but omit execution interfaces, detailed scoring, or other sections.
The manifest records exact character ranges and how the retained passages join.
The added `ai_research` and `systems_engineering` categories describe the selected focus, not upstream labels or expected tiers.
Question IDs use `researchrubrics/` or `frontiercs/` prefixes followed by the original source ID.

Research topics include experimental validation, instance segmentation, predictive coding, model training, GPU kernels, and symbolic regression.
Engineering topics include interruptible scheduling, cloud broadcast, matrix multiplication, LLM routing, query optimization, simulation, compaction, and vector retrieval.
Selection used task requirements and topic coverage, without picker predictions.
The user labeled two supplemental requests `tier_2` and 18 requests `tier_3`.
These labels express the intended routing policy, not measured model quality.
Label each prompt as displayed in the CSV, without assuming its source requires a particular tier.

Report the original 80 and the advanced supplement separately, alongside any combined score.
The supplement deliberately changes the difficulty distribution and does not represent typical user traffic.
Record stakes and tool assumptions in `notes`. A low-stakes task can still require difficult reasoning.

The original FrontierCS statements refer to benchmark datasets, repositories, or hardware, which are not bundled here.
This CSV supports request classification, not execution of the benchmark tasks.
The 16 FrontierCS requests use explicit excerpts because the full specifications exceed the current prompt budget.
The actual classifier tokenizer counts 27-329 tokens per added prompt. All 20 fit the normal `config.yaml` policy's 426-token prompt budget.
The manifest records the configuration and tokenizer hashes for that measurement.
A subsequent audit of all 100 requests found no truncation under the current configuration.
The picker keeps only leading prompt tokens within its 512-token total input budget, including the question and tier descriptions.
After a configuration change, measure the available prompt budget again. Report truncated requests separately.
Do not shorten requests after observing predictions.

## Creative-writing suite

[Cases](suites/creative-writing.jsonl) include prose, poetry, screenplays, novels, plot development, character design, and fan fiction.
All English rows in these categories are included; no predictions influenced selection.
The source revision, selection rules, counts, and file hashes are in the [manifest](suites/creative-writing-source.json).

Before inference, review the category-to-tier labels and freeze the candidate, configuration, cases, and scoring policy.
The release pass/fail threshold is not yet agreed. No evaluation results are available.

Report writing-tier recall, the distribution of other selected tiers, HTTP failures, and invalid model membership.
Report cases that exceed the actual tokenizer budget separately. The service silently keeps only leading prompt tokens within its 512-token total input budget.
Do not rewrite long queries to improve results. Cross-tier or unclear cases need separate reporting after label review.
A writing-only suite measures recall; it cannot establish writing precision or overall picker accuracy.

## Attribution

MT-Bench prompts are from *Judging LLM-as-a-Judge with MT-Bench and Chatbot Arena*, Lianmin Zheng and colleagues (2023),
[paper](https://arxiv.org/abs/2306.05685), [source at the pinned revision](https://github.com/lm-sys/FastChat/blob/b494d0c6b4e7935f1764f8439e75da3e66beccc7/fastchat/llm_judge/data/mt_bench/question.jsonl).
The source repository includes an [Apache-2.0 license](https://github.com/lm-sys/FastChat/blob/b494d0c6b4e7935f1764f8439e75da3e66beccc7/LICENSE).
A copy is included at [third_party/Apache-2.0.txt](../third_party/Apache-2.0.txt).
We selected the first turn of each question and added blank `expected_tier` and `notes` columns. Prompt text is unchanged.

ResearchRubrics prompts are from *ResearchRubrics: A Benchmark of Prompts and Rubrics For Evaluating Deep Research Agents*,
[paper](https://arxiv.org/abs/2511.07685), [pinned dataset](https://huggingface.co/datasets/ScaleAI/researchrubrics/tree/85de3115053d1453ed612caacf4a405edc1ad756).
The dataset card declares MIT terms. The project [license copy](../third_party/researchrubrics/LICENSE) retains its original copyright notice.

FrontierCS statements are from *FrontierCS: Evolving Challenges for Evolving Intelligence*, Qiuyang Mang and colleagues,
[paper](https://arxiv.org/abs/2512.15699), [pinned dataset](https://huggingface.co/datasets/FrontierCS/Frontier-CS/tree/2dbadf74ad06015c842444f77b39356a55576987).
The pinned dataset card declares Apache-2.0 terms, while its [LICENSE file](../third_party/frontiercs/LICENSE) contains MIT terms.
Both texts are retained, including the [Apache-2.0 copy](../third_party/Apache-2.0.txt), without changing the source notices.
The source manifest identifies the selected statements, including two tasks in the dataset's `2.0` category.
The supplement contains four unchanged prompts and 16 documented excerpts, with IDs, focus categories, and manual annotation columns.

Creative prompts are from *WritingBench: A Comprehensive Benchmark for Generative Writing*, Yuning Wu and colleagues (2025),
[paper](https://arxiv.org/abs/2503.05244), [source at the pinned revision](https://github.com/X-PLUG/WritingBench/tree/ae2d5176449b7b769815482641d35926f26793eb).
They retain the upstream [Apache-2.0 license](https://github.com/X-PLUG/WritingBench/blob/ae2d5176449b7b769815482641d35926f26793eb/LICENSE);
a copy is included at [third_party/Apache-2.0.txt](../third_party/Apache-2.0.txt).
We selected English creative categories and added suite IDs, expected tiers, provenance, and mapping rationales. Prompt text is unchanged.
These derived selection results will not be official WritingBench generation scores.

WildBench declares CC BY 4.0; its research/educational-use and Responsible Use Guidelines statements still need review before importing cases.
CLINC150's Hugging Face card declares CC BY 3.0. Imported cases must preserve their separate dataset attribution and license terms.
