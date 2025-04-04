PROTO_NAMES=(
    "auth"
    "conversation"
    "errinfo"
    "relation"
    "group"
    "jssdk"
    "msg"
    "msggateway"
    "push"
    "rtc"
    "sdkws"
    "third"
    "user"
    "statistics"
    "wrapperspb"
)
ddir=DestinationPath
sdir=SourcePath
for name in "${PROTO_NAMES[@]}"; do
    cat ${sdir}/${name}/${name}.proto > ${ddir}/${name}/${name}.proto
    if [ $? -ne 0 ]; then
        echo "error when processing ${name}"
        exit #?
    fi    
done
