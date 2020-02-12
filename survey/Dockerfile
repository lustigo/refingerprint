FROM node:13-buster-slim AS frontend
WORKDIR /app
COPY frontend/package* ./
RUN ["npm","install"]
COPY frontend/tsconfig.json tsconfig.json
COPY frontend/.eslintrc.json .eslintrc.json
COPY frontend/babel.config.js babel.config.js
COPY frontend/vue.config.js vue.config.js
COPY frontend/public public/
COPY frontend/src src/
RUN ["npm","run","build"]

FROM node:13-buster-slim AS backend
WORKDIR /app
COPY backend/package* ./
RUN ["npm","install"]
COPY backend/tsconfig.json tsconfig.json
COPY backend/.eslintrc.json .eslintrc.json
COPY backend/src src/
RUN ["npm","run","build"]

FROM node:13-buster-slim
WORKDIR /app
COPY --from=frontend /app/dist frontend/
COPY --from=backend /app/dist ./
COPY --from=backend /app/package* ./
RUN ["npm","install","--production"]
ENV APP_PATH /app/frontend/
ENV DATA_PATH /data/
ENV PORT 443
ENV TLS_KEY_PATH /data/key.pem
ENV TLS_CERT_PATH /data/crt.pem
ENV RECAPTCHA_SECRET PLACEHOLDER
EXPOSE 443
ENTRYPOINT [ "node","index.js" ]
