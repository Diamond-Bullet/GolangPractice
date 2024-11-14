# .kinit_auto.sh, kinit automatically when powering on.
fail=false
for ((i = 0; i < 5; i++)); do
  kdestroy && kinit --password-file=/root/.kinit_password -l 86400 root@666.com
  if [ "$(klist | wc -l)" -ge 4 ]; then
    break
  fi
  fail=true
  echo "$(date "+%Y-%m-%d %H:%M:%S"), kinit failed in ${i}" >>/root/.kinit_log.txt
  sleep 1s
done
if [ $fail == true ]; then
  echo >>/root/.kinit_log.txt
fi
