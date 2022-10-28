FROM golang:1.18.7-alpine

ENV USER ariefid
ENV LOGINSHELL /bin/bash
ENV HOME /home/${USER}
ENV WORKDIR ${HOME}/apps/echo
ENV UID 1048
ENV GID 1048
ENV MYSQL_HOST ariefid-db-mysql
ENV MYSQL_PORT 3306
ENV MYSQL_USER docker
ENV MYSQL_PASSWORD docker
ENV MYSQL_DBNAME todo
ENV MYSQL_ROOT_PASSWORD rootAbraCadabra
ENV ECHO_DOCKER_PORT 3030
ENV ECHO_EXPOSE_PORT 3030
ENV ECHO_BUILD_OUTPUT ariefid-devcode-golang-echo

RUN addgroup \
    --gid "${GID}" \
    "${USER}"

RUN adduser \
    --home "${HOME}" \
    --gecos "" \
    --shell "${LOGINSHELL}" \
    --ingroup "${USER}" \
    --disabled-password \
    --uid "${UID}" \
    "${USER}"

USER ${USER}

RUN mkdir -p ${WORKDIR}

WORKDIR ${WORKDIR}

COPY --chown=${USER}:${USER} ./apps/echo .

RUN go build -o ${ECHO_BUILD_OUTPUT}

RUN chmod 755 ./entrypoint.sh

RUN chmod 755 ./${ECHO_BUILD_OUTPUT}

CMD [ "sh", "-c", "--", "./entrypoint.sh" ]
