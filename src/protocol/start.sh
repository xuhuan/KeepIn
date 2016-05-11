#!/bin/sh
GMOCK=gmock-1.7.0.zip
PROTOBUF=protobuf-2.6.1
CUR_DIR=
SRC_DIR=./src
DST_DIR=./_gen

init_protobuf(){
    cd protobuf
    tar -xf $PROTOBUF.tar.gz
    cd $PROTOBUF
    ./configure --prefix=$CUR_DIR/protobuf
    make
    make install
    export PATH=$CUR_DIR/protobuf:$PATH
}

check_user() {
    if [ $(id -u) != "0" ]; then
        echo "Error: You must be root to run this script"
        exit 1
    fi
}

get_cur_dir() {
    # Get the fully qualified path to the script
    case $0 in
        /*)
            SCRIPT="$0"
            ;;
        *)
            PWD_DIR=$(pwd);
            SCRIPT="${PWD_DIR}/$0"
            ;;
    esac
    # Resolve the true real path without any sym links.
    CHANGED=true
    while [ "X$CHANGED" != "X" ]
    do
        # Change spaces to ":" so the tokens can be parsed.
        SAFESCRIPT=`echo $SCRIPT | sed -e 's; ;:;g'`
        # Get the real path to this script, resolving any symbolic links
        TOKENS=`echo $SAFESCRIPT | sed -e 's;/; ;g'`
        REALPATH=
        for C in $TOKENS; do
            # Change any ":" in the token back to a space.
            C=`echo $C | sed -e 's;:; ;g'`
            REALPATH="$REALPATH/$C"
            # If REALPATH is a sym link, resolve it.  Loop for nested links.
            while [ -h "$REALPATH" ] ; do
                LS="`ls -ld "$REALPATH"`"
                LINK="`expr "$LS" : '.*-> \(.*\)$'`"
                if expr "$LINK" : '/.*' > /dev/null; then
                    # LINK is absolute.
                    REALPATH="$LINK"
                else
                    # LINK is relative.
                    REALPATH="`dirname "$REALPATH"`""/$LINK"
                fi
            done
        done

        if [ "$REALPATH" = "$SCRIPT" ]
        then
            CHANGED=""
        else
            SCRIPT="$REALPATH"
        fi
    done
    # Change the current directory to the location of the script
    CUR_DIR=$(dirname "${REALPATH}")
}

build_protobuf(){    
	rm -rf DST_DIR	

	#C++
	mkdir -p $DST_DIR/cpp
	protoc -I=$SRC_DIR --cpp_out=$DST_DIR/cpp/ $SRC_DIR/*.proto

	#JAVA
	mkdir -p $DST_DIR/java
	protoc -I=$SRC_DIR --java_out=$DST_DIR/java/ $SRC_DIR/*.proto

	#PYTHON
	mkdir -p $DST_DIR/python
	protoc -I=$SRC_DIR --python_out=$DST_DIR/python/ $SRC_DIR/*.proto

	#JS
	mkdir -p $DST_DIR/js
	protoc -I=$SRC_DIR --js_out=import_style=commonjs,binary:$DST_DIR/js/ $SRC_DIR/*.proto

	#go
	mkdir -p $DST_DIR/go
	protoc -I=$SRC_DIR --go_out=$DST_DIR/go/ $SRC_DIR/*.proto
    # protoc -I=./src --go_out=./gen/go/ ./src/*.proto
	# for i in $SRC_DIR/*.proto; do
	# 	protoc -I=$SRC_DIR --go_out=$DST_DIR/go/ $i
	# done
}

deploy_lib(){
    mkdir -p $2/lib/linux/
    cp protobuf/lib/libprotobuf-lite.a $2/lib/linux/
    cp  -r protobuf/include/* $2/
}

deploy_protobuf(){

	#C++
	cp $DST_DIR/cpp/* $2/

	#JAVA
	cp $DST_DIR/java/* $2/

	#PYTHON
	cp $DST_DIR/python/* $2/

	#JS
	cp $DST_DIR/js/* $2/

	#go
	cp $DST_DIR/go/* $2/
}

deploy_go_protobuf(){
	#go
	cp $DST_DIR/go/* ../../protocol/
}

print_help() {
	echo "Usage: "
	echo "  $0 init ---  install protobuf"
	echo "  $0 build --- build all"
	echo "  $0 deploy_lib floder --- deploy lib to floder"
	echo "  $0 deploy_go floder --- deploy go to floder"
	echo "  $0 deploy floder --- deploy to floder"
}

case $1 in
	init)
		check_user
		get_cur_dir
		init_protobuf
		;;
	build)
		build_protobuf
		;;
	deploy)
		if [ $# != 2 ]; then 
			echo $#
			print_help
			exit
		fi

		echo $2
		echo "deploy..."
		deploy_protobuf $2
		;;
	deploy_lib)
		if [ $# != 2 ]; then 
			echo $#
			print_help
			exit
		fi

		echo $2
		echo "deploy lib..."
		deploy_lib $2
		;;
	deploy_go)
		echo "deploy go..."
		deploy_go_protobuf
		;;
	*)
		print_help
		;;
esac