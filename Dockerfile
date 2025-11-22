# =========================
# 1. Build SvelteKit frontend
# =========================
FROM node:22-alpine AS frontend

# Avoid interactive prompts and enable pnpm via Corepack
ENV COREPACK_ENABLE_DOWNLOAD_PROMPT=0
RUN npm install -g corepack@latest && corepack enable pnpm

WORKDIR /app

COPY server/api ./server/api

WORKDIR /app/web

# Copy dependency files and install
COPY web/pnpm-lock.yaml web/package.json ./
RUN pnpm install --frozen-lockfile

# Copy the rest of the frontend source and build
COPY web/ .

RUN pnpm run gen:api
RUN pnpm build

# =========================
# 2. Build Go backend
# =========================
FROM golang:1.25-alpine AS backend

WORKDIR /app

# Copy Go mod files and download dependencies
COPY server/go.mod server/go.sum ./
RUN go mod download

# Copy backend source code
COPY server/ .

# Build Go binary
ARG TARGETOS
ARG TARGETARCH
ARG VERSION

RUN CGO_ENABLED=0 GOOS=$TARGETOS GOARCH=$TARGETARCH \
    go build -ldflags="-s -w -X pycrs.cz/what-it-doo/pkg/version.Version=${VERSION}" \
    -o /app/bin/server ./cmd/api

RUN CGO_ENABLED=0 GOOS=$TARGETOS GOARCH=$TARGETARCH \
    go build -ldflags="-s -w -X pycrs.cz/what-it-doo/pkg/version.Version=${VERSION}" \
    -o /app/bin/worker ./cmd/worker

# =========================
# 3. Final minimal image
# =========================
FROM alpine:latest AS production

WORKDIR /app
COPY --from=backend /app/bin/server ./bin/server
COPY --from=backend /app/bin/worker ./bin/worker
COPY docker-entrypoint.sh ./bin/entrypoint.sh 
COPY --from=frontend /app/web/build ./web

EXPOSE 8080
ENTRYPOINT ["/app/bin/entrypoint.sh"]
