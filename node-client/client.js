const grpc = require("@grpc/grpc-js");
const protoLoader = require("@grpc/proto-loader");
const path = require("path");

const PROTO_PATH = path.join(__dirname, "..", "user.proto");

const packageDefinition = protoLoader.loadSync(PROTO_PATH, {
    keepCase: false,
    longs: String,
    enums: String,
    defaults: true,
    oneofs: true
});

const proto = grpc.loadPackageDefinition(packageDefinition);

const client = new proto.user.UserService(
    "localhost:50051",
    grpc.credentials.createInsecure()
);

client.getUserProfile({ userId: 1 }, (error, response) => {
    if (error) {
        console.error("Terjadi kesalahan:", error.message);
        client.close();
        return;
    }

    console.log("=== PROFIL PENGGUNA ===");
    console.log("Nama   :", response.name);
    console.log("Email  :", response.email);
    console.log("Telepon:", response.phone);
    
    client.close();
});