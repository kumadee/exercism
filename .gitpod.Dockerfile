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
    mkdir -p ~/.config/exercism && \
    cat ./shell/exercism_completion.bash > /home/gitpod/.bashrc.d/10-exercism && \
    rm -rf exercism-linux-64bit.tgz LICENSE shell
#
# More information: https://www.gitpod.io/docs/config-docker/
