FROM debian:bookworm-slim

RUN apt update && apt install -y --no-install-recommends \
  imagemagick \
  ghostscript \
  ca-certificates \
  file \
  uuid-runtime \
  && rm -rf /var/lib/apt/lists/*

RUN sed -i 's/policy domain="coder" rights="none" pattern="PDF"/policy domain="coder" rights="read|write" pattern="PDF"/' /etc/ImageMagick-6/policy.xml && \
  sed -i 's/policy domain="resource" name="memory" value="256MiB"/policy domain="resource" name="memory" value="2GiB"/' /etc/ImageMagick-6/policy.xml && \
  sed -i 's/policy domain="resource" name="map" value="512MiB"/policy domain="resource" name="map" value="4GiB"/' /etc/ImageMagick-6/policy.xml && \
  sed -i 's/policy domain="resource" name="width" value="16KP"/policy domain="resource" name="width" value="128KP"/' /etc/ImageMagick-6/policy.xml && \
  sed -i 's/policy domain="resource" name="height" value="16KP"/policy domain="resource" name="height" value="128KP"/' /etc/ImageMagick-6/policy.xml && \
  sed -i 's/policy domain="resource" name="disk" value="1GiB"/policy domain="resource" name="disk" value="8GiB"/' /etc/ImageMagick-6/policy.xml

WORKDIR /app

COPY bin/pdf-fix .

COPY scripts/ ./scripts/
COPY static/ ./static/
COPY templates/ ./templates/

RUN mkdir -p pdfTmp uploads && \
  chmod 777 pdfTmp uploads

HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
  CMD curl -f http://localhost:8080/health || exit 1

ENV GIN_MODE=release

EXPOSE 8080

CMD ["./pdf-fix"]
