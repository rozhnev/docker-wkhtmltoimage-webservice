FROM ubuntu:20.04

RUN apt-get update && apt-get install -y wget libappindicator1 fonts-liberation libasound2 libatk-bridge2.0-0 libcurl4 \
    libdrm2 libgbm1 libnspr4 libnss3 libxkbcommon0 xdg-utils
RUN apt-get install -f

RUN wget https://dl.google.com/linux/direct/google-chrome-stable_current_amd64.deb
# RUN dpkg -i google-chrome-stable_current_amd64.deb

RUN apt --fix-broken install
RUN dpkg -i google-chrome-stable_current_amd64.deb

RUN wget https://go.dev/dl/go1.17.5.linux-amd64.tar.gz
RUN tar -C /usr/local -xzf go1.17.5.linux-amd64.tar.gz
RUN echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.profile 
# RUN source ~/.profile

# Move to working directory /build
WORKDIR /build

# Copy the code into the container
COPY *.go ./
COPY snapshots /tmp/snapshots/


# # Build the application
ENV PATH=$PATH:/usr/local/go/bin \
    GO111MODULE=off \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

RUN go build -o main .

# # Move to /dist directory as the place for resulting binary folder
WORKDIR /dist

# # Copy binary from build to main folder
RUN cp /build/main .

# # Export necessary port
EXPOSE 80

# # Command to run when starting the container
CMD ["/dist/main"]
