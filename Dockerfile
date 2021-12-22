FROM ubuntu:20.04

RUN apt-get update && apt-get install -y wget libexpat1 ucf fonts-dejavu-core libjpeg-turbo8 \
    libx11-data libx11-6 libxau6 libbsd0 libxdmcp6 libxcb1 fontconfig-config libfontconfig1  libpng16-16 libfreetype6 fontconfig\
    libxext6 libxrender1  xfonts-utils xfonts-75dpi xfonts-base libfontenc1 x11-common xfonts-encodings \
    python3 pip

RUN wget https://github.com/wkhtmltopdf/packaging/releases/download/0.12.6-1/wkhtmltox_0.12.6-1.focal_amd64.deb

RUN dpkg -i wkhtmltox_0.12.6-1.focal_amd64.deb


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
