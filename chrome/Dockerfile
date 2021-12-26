FROM ubuntu:20.04

ARG DEBIAN_FRONTEND=noninteractive
ENV TZ=Europe/Amsterdam
RUN apt-get update && apt-get install -y tzdata

RUN apt-get install -y -qq wget libappindicator1 fonts-liberation libasound2 libatk-bridge2.0-0 libcurl4 \
    libdrm2 libgbm1 libnspr4 libnss3 libxkbcommon0 xdg-utils  libgtk-3-0 \
    libgtk-3-0 libcolord2 liblcms2-2 libepoxy0 libjson-glib-1.0-0 librest-0.7-0 libwayland-cursor0 \ 
    libwayland-egl1 libgtk-3-common libjson-glib-1.0-common libsoup-gnome2.4-1  libsoup2.4-1 dconf-gsettings-backend \
    dconf-service libdconf1 default-dbus-session-bus libpam-systemd  systemd libcap2 libcryptsetup12 libip4tc2 libkmod2 \
    libargon2-1 libdevmapper1.02.1  libjson-c4 systemd-timesyncd systemd-sysv  glib-networking libproxy1v5 glib-networking-services \
    glib-networking-common gsettings-desktop-schemas

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
RUN mkdir /tmp/snapshots/


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
