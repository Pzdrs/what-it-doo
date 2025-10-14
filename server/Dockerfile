# =========================
# 1. Build SvelteKit frontend
# =========================
FROM node:22-alpine AS frontend

# Avoid interactive prompts and enable pnpm via Corepack
ENV COREPACK_ENABLE_DOWNLOAD_PROMPT=0
RUN npm install -g corepack@latest && corepack enable pnpm

WORKDIR /app

# Copy dependency files and install
COPY web/pnpm-lock.yaml web/package.json ./
RUN pnpm install --frozen-lockfile

# Copy the rest of the frontend source and build
COPY web/ .
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

# Copy built frontend files into Go’s static directory

# Build Go binary
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/bin/server ./cmd/api

# =========================
# 3. Final minimal image
# =========================
FROM gcr.io/distroless/base-debian12 AS production

WORKDIR /app
COPY --from=backend /app/bin/server ./bin/server
COPY --from=frontend /app/build ./static

EXPOSE 8080
ENTRYPOINT ["/app/bin/server"]
