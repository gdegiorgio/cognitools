
#!/bin/bash
COGNITOOLS_HOME=$HOME/.cognitools




install(){


    if [ -f $COGNITOOLS_HOME/bin/cognitools ]; then
        echo "Cognitools binary already exists in $COGNITOOLS_HOME/bin/cognitools"
        return
    fi

    
    ARCH=""
    MACHINE="$(uname -m)"

    if [ "$MACHINE" = "x86_64" ]; then
        ARCH="amd64"
    elif [ "$MACHINE" = "aarch64" ] || [ "$MACHINE" = "arm64" ]; then
        ARCH="arm64"
    else
        echo "Unsupported architecture: $MACHINE"
        return 1
    fi

    OS_NAME="$(uname -s)"
    if [ "$OS_NAME" = "Darwin" ]; then
        OS="darwin"
    elif [ "$OS_NAME" = "Linux" ]; then
        OS="linux"
    else
        echo "Unsupported OS: $OS_NAME"
        return 1
    fi

    echo "Installing cognitools for $OS $ARCH"

    mkdir -p $COGNITOOLS_HOME/bin

    cd $COGNITOOLS_HOME

    curl -L -o "cognitools_${OS}_${ARCH}.tar.gz" "https://github.com/gdegiorgio/cognitools/releases/latest/download/cognitools_${OS}_${ARCH}.tar.gz"

    tar -xzf "cognitools_${OS}_${ARCH}.tar.gz"
    rm "cognitools_${OS}_${ARCH}.tar.gz"

    mv cognitools $COGNITOOLS_HOME/bin/cognitools
    chmod +x $COGNITOOLS_HOME/bin/cognitools
    


    if [ $SHELL == "/bin/bash" ]; then
        echo "export COGNITOOLS_HOME=$COGNITOOLS_HOME" >> $HOME/.bashrc
        echo 'export PATH=$PATH:$COGNITOOLS_HOME/bin' >> $HOME/.bashrc    
        source_line="source $HOME/.bashrc"
    elif [ $SHELL == "/bin/zsh" ]; then
        echo "export COGNITOOLS_HOME=$COGNITOOLS_HOME" >> $HOME/.zshrc
        echo 'export PATH=$PATH:$COGNITOOLS_HOME/bin' >> $HOME/.zshrc
        source_line="source $HOME/.zshrc"
    elif [ $SHELL == "/bin/sh" ]; then
        echo "export COGNITOOLS_HOME=$COGNITOOLS_HOME" >> $HOME/.profile
        echo 'export PATH=$PATH:$COGNITOOLS_HOME/bin' >> $HOME/.profile
        source_line="source $HOME/.profile"
    else
        echo "Unsupported shell"
        return
    fi


    echo "Cognitools installed successfully, restart your shell or run $source_line to use it"

}



install
