const PROTO_PATH = "../pb/notification.proto"

const grpc = require("grpc")
var protoLoader = require("@grpc/proto-loader");

var packageDefinition = protoLoader.loadSync(PROTO_PATH, {
    keepCase: true,
    longs: String,
    enums: String,
    arrays: true
});

var NotificationService = grpc.loadPackageDefinition(packageDefinition).NotificationService;

const client = new NotificationService(
    "localhost:30043",
    grpc.credentials.createInsecure()
);

module.exports = client;