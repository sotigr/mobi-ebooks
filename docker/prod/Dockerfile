FROM golang:1.22 AS builder
    
COPY cmd /src/cmd   
COPY internal /src/internal
COPY static /src/static
COPY templates /src/templates

COPY go.mod /src/go.mod 
COPY go.sum /src/go.sum
  
WORKDIR /src 

RUN go build -o dist ./cmd

FROM debian:bookworm-slim

ENV DEBIAN_FRONTEND=noninteractive

RUN apt update
RUN apt update && apt install -y calibre && rm -rf /var/lib/apt/lists/*

COPY --from=builder /src/dist /dist/dist
COPY --from=builder /src/static/ /dist/static/ 
COPY --from=builder /src/templates/ /dist/templates/ 

WORKDIR /dist 
 
CMD [ "/dist/dist" ] 