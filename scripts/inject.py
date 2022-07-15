import sys
import os



def main(dirname, endpoint):
    for filename in os.listdir(dirname):

        file_path = os.path.join(dirname, filename)
        print("curl -X POST --data-binary @" + file_path + " " + endpoint)
        os.system("curl -X POST --data-binary @" + file_path + " " + endpoint + " > /dev/null")


    #curl -X POST --data-binary @/home/emurray/Downloads/Morning_Run.fit http://localhost:8080/api/v1/activities/upload


if __name__=="__main__":
    if len(sys.argv) != 3:
        print("Usage: python3 inject.py <name-of-directory> <endpoint>")
        sys.exit(1)
    dirname = sys.argv[1]
    endpoint = sys.argv[2]
    main(dirname, endpoint)
