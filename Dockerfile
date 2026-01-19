FROM debian:bookworm-slim

RUN apt update && apt install -y --no-install-recommends \
  imagemagick \
  ghostscript \
  ca-certificates \
  && rm -rf /var/lib/apt/lists/*

# RUN sed -i 's/policy domain="coder" rights="none" pattern="PDF"/policy domain="coder" rights="read|write" pattern="PDF"/' /etc/ImageMagick-6/policy.xml

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
