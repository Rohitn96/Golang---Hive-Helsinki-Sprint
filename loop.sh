count=$1

if [ "$count" -gt 100 ]; then
	count=100
fi

for ((i=1; i<=count; i++))
do
	echo "This is loop number $i"
done
