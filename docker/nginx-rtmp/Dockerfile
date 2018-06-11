FROM tiangolo/nginx-rtmp

# Install ffmpeg
RUN apt-get update && \
    apt-get install -y --no-install-recommends ffmpeg

# Create the video and thumbnail output directories
RUN mkdir -p /var/www/live && \
    mkdir -p /var/www/record && \
    mkdir -p /var/www/thumbnail

# Assign correct user and group to newly created directories
RUN chown www-data:www-data /var/www/live && \
    chown www-data:www-data /var/www/record && \
    chown www-data:www-data /var/www/thumbnail

# Create some Nginx config directories
RUN mkdir -p /etc/nginx/sites-available && \
    mkdir -p /etc/nginx/sites-enabled

# Copy over the Nginx config files
COPY nginx.conf /etc/nginx/nginx.conf
COPY rtmp.conf /etc/nginx/rtmp.conf
COPY live.conf /etc/nginx/sites-available/live.conf

# Enable live.conf
RUN ln -s /etc/nginx/sites-available/live.conf /etc/nginx/sites-enabled/live.conf

# Make sure our live logs are exposed by Docker
RUN ln -sf /dev/stdout /var/log/nginx/live.access.log && \
    ln -sf /dev/stderr /var/log/nginx/live.error.log

EXPOSE 80
