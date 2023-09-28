# ES query
curl -H "Content-Type: application/json" -u elastic:gRYppQtdTpP3nFzH -XPOST 'http://172.16.5.146:9200/imissu/_doc/_search?pretty' -d '{
  "query": {
            "bool": {
              "should": [
                { "term": { "user_id": "10" } },
                { "term": { "stage_name": "10" } }
              ]
            }
  },
  "sort": {
    "_script": {
      "type": "number",
      "script": {
        "lang": "painless",
        "source": " if (doc.uk.value.contains(params.q)) { if (doc.sk.value.contains(params.q)) {return Math.min(doc.sk.value.length()*10+1,
doc.uk.value.length()*10);} else {return doc.uk.value.length()*10;}} else {return doc.sk.value.length()*10+1;}",
        "params": {
          "q": "10"
        }
      },
      "order": "asc"
    }
  }
}'

# proto generation
protoc --micro_out=. --go_out=. ./customer.proto


####################### Mac OS #################
# Mac下使用了zsh会不执行/etc/profile文件，当然，如果用原始的是会执行。
# 转而执行的是这两个文件，每次登陆都会执行：~/.zshrc与/etc/zshenv、/etc/zshrc

# clear screen
cmd + k

# list processes
ps -lef

# show all files' size in current directory
du -sh *