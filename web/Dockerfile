FROM nginx:alpine

# ลบ config เดิม
RUN rm /etc/nginx/conf.d/default.conf

# ใส่ config ของเรา
COPY nginx.conf /etc/nginx/conf.d/default.conf

# ใส่ไฟล์ Flutter Web
COPY build/ /usr/share/nginx/html

EXPOSE 8088
