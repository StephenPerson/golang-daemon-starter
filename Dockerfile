FROM busybox
WORKDIR /golang-daemon-starter
COPY . .
EXPOSE 8080
ENTRYPOINT ["./golang-daemon-starter"]
CMD ["console"]