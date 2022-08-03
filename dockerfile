FROM golang:1.18

# create directory app
RUN mkdir /app

# set or make /app our working directory
WORKDIR /app

# copy all files to /app
COPY . .

RUN go build -o g4-api

CMD [ "./g4-api" ]
