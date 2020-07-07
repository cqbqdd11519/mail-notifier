FROM golang:1.14-buster AS build

COPY ../.. /source
RUN make -C /source build

FROM scratch

COPY --from=build /source/bin/mail-notifier /mail-notifier

ENTRYPOINT ["/mail-notifier"]
