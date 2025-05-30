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

for name in "${PROTO_NAMES[@]}"; do
    rm -r ./../${name}
    if [ $? -ne 0 ]; then
        echo "error when processing ${name}"
        exit #?
    fi    
done
