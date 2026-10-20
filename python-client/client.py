import grpc

import user_pb2
import user_pb2_grpc


def main():
    with grpc.insecure_channel("localhost:50051") as channel:
        stub = user_pb2_grpc.UserServiceStub(channel)

        try:
            response = stub.GetUserProfile(
                user_pb2.UserRequest(user_id=1)
            )

            print("=== PROFIL PENGGUNA ===")
            print(f"Nama  : {response.name}")
            print(f"Email : {response.email}")
            print(f"Telepon: {response.phone}")
            
        except grpc.RpcError as error:
            print(f"Terjadi kesalahan: {error.code()}")
            print(error.details())


if __name__ == "__main__":
    main()