FROM gitpod/workspace-full

USER gitpod

# Install custom tools, runtime, etc. using apt-get
#
# RUN sudo apt-get -q update && \
#    sudo apt-get install -yq bastet && \
#    sudo rm -rf /var/lib/apt/lists/*
RUN wget https://github.com/exercism/cli/releases/download/v3.0.13/exercism-linux-64bit.tgz && \
    tar -xf exercism-linux-64bit.tgz && \
    mv exercism /usr/local/bin && \
    chmod a+rx /usr/local/bin/exercism && \
    rm exercism-linux-64bit.tgz
#
# More information: https://www.gitpod.io/docs/config-docker/
