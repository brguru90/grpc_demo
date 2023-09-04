const PROTO_PATH = "../pb/notification.proto"

const grpc = require("grpc")
var protoLoader = require("@grpc/proto-loader");
const { v4: uuidv4 } = require("uuid");

var packageDefinition = protoLoader.loadSync(PROTO_PATH, {
    keepCase: true,
    longs: String,
    enums: String,
    arrays: true
});

var notificationsProto = grpc.loadPackageDefinition(packageDefinition);
const server = new grpc.Server();

const notifications=require("./db")

server.addService(notificationsProto.NotificationService.service,{
    GetAllNotifications:(_, callback)=>{
        if(!notifications.length){
            callback({
                code: grpc.status.NOT_FOUND,
                details: "Not found"
            })
            return
        }
        callback(null,{notification_list:notifications})

    },
    AddNotifications:(call, callback)=>{
        const {title,content}=call.request

        notifications.push({
            id:uuidv4(),
            time_stamp:new Date().getTime(),
            title,
            content,
        })

        callback(null,{notification_list:notifications})

    },
    RemoveNotifications:(call, callback)=>{
        const {id}=call.request
        const index=notifications.findIndex(n=>n.id==id)
        if(index==-1){
            callback({
                code: grpc.status.NOT_FOUND,
                details: "Not found"
            })
            return
        }
        notifications.splice(index,1)
        callback(null,{})
    }
})



server.bind("127.0.0.1:30043", grpc.ServerCredentials.createInsecure());
console.log("Server running at http://127.0.0.1:30043");
server.start();
