file_count=$(find . -type f | wc -l)
folder_count=$(find . -type d | wc -l)

total_count=$((file_count + folder_count))

echo $total_count