# Record the start time in nanoseconds
start_time=$(date +%s%N)

# Your command or script
curl -v -X POST http://10.10.144.145:8081/var/api/v1/calc/all -H 'Content-Type: application/json' -d @files/request.txt

# Record the end time in nanoseconds
end_time=$(date +%s%N)

# Calculate the duration in nanoseconds
duration=$((end_time - start_time))

# Convert the duration to seconds
formatted_duration=$((duration/1000000))

# Print the duration
echo -e "\nTime taken: $formatted_duration ms"