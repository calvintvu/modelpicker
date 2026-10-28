FROM python:3.13.18-slim-bookworm@sha256:93498f36dc7e24ef233bfdf6cc17d4b8bd32d94ee737143e906460d5f5a914c5

COPY docker/model-download-requirements.txt /tmp/requirements.txt
RUN pip install --no-cache-dir --require-hashes --root-user-action=ignore -r /tmp/requirements.txt

COPY LICENSE THIRD_PARTY_NOTICES.md /usr/share/licenses/ez-model-router/
COPY third_party/ /usr/share/licenses/ez-model-router/third_party/

# Public model downloads need no credentials or persistent global cache.
ENV HF_HOME=/tmp/huggingface \
    HF_HUB_DISABLE_IMPLICIT_TOKEN=1 \
    HF_HUB_DISABLE_TELEMETRY=1

ENTRYPOINT ["hf"]
