date +%s
date
start_time=$(date +%s)
./custom-metrics
date
date +%s
end_time=$(date +%s)
sleep 300
python3 ~/fetch_metrics.py $start_time $end_time 200

