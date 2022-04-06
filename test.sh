echo "Testing locations API"
# # curl -X GET https://schematics.cloud.ibm.com/v1/locations -H "Authorization: <iam_token>"
# url=https://schematics.cloud.ibm.com/v1/locations

# # url=https://temptemp3.github.io
# # just some url
 
# curl ${url} -I -o headers -s
# # download file
 
# cat  headers
# # response headers
# ## expect
# #HTTP/2 200
# #server: GitHub.com
# #content-type: text/html; charset=utf-8
# #strict-transport-security: max-age=31557600
# #last-modified: Thu, 03 May 2018 02:30:03 GMT
# #etag: "5aea742b-e12"
# #access-control-allow-origin: *
# #expires: Fri, 25 Jan 2019 23:07:17 GMT
# #cache-control: max-age=600
# #x-github-request-id: 8808:5B91:2A4802:2F2ADE:5C4B944C
# #accept-ranges: bytes
# #date: Fri, 25 Jan 2019 23:12:37 GMT
# #via: 1.1 varnish
# #age: 198
# #x-served-by: cache-nrt6148-NRT
# #x-cache: HIT
# #x-cache-hits: 1
# #x-timer: S1548457958.868588,VS0,VE0
# #vary: Accept-Encoding
# #x-fastly-request-id: b78ff4a19fdf621917cb6160b422d6a7155693a9
# #content-length: 3602
 
# cat headers | head -n 1 | cut '-d ' '-f2'
# # get response code
# ## expect
# #200