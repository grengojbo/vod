vod
===

Video on Demand 

кодируем
High profile H264 — 3 качества (X1.mp4 X2.mp4 X3.mp4)
Baseline H.264 — 2 качества (X4.mp4 X5.mp4)

под flash — отдаем OpenHttpStreamer'ом X1.mp4 X2.mp4 X3.mp4
под iphone — ErlyVideo X4.mp4 X5.mp4
под iPad — ErlyVideo X1.mp4 X2.mp4 X3.mp4 (там мощный чип декодирования и экран большой)
под android phone — http progressive download c ручным переключением X4.mp4 X5.mp4
под android pad — http progressive download c ручным переключением X1.mp4 X2.mp4 X3.mp4

MPEG-DASH GStreamer
gst-launch-1.0 -v -m v4l2src ! video/x-raw,width=640,height=480 \
        ! x264enc tune=zerolatency ! "video/x-h264,profile=baseline" ! h264parse \
        ! flvmux streamable=true name=mux \
        alsasrc ! queue ! audioconvert ! voaacenc ! aacparse ! mux. \
        mux. ! rtmpsink location=rtmp://localhost/myapp/mystream

И еще трансляцию разрезаю по битрейтам — вот пример
ffmpeg -i "rtsp://10.1.2.71/play1.sdp" -threads 2 -vcodec libx264 -preset UltraFast -rtbufsize 10000k -analyzeduration 0 -tune zerolatency -s 640x480 -acodec libmp3lame -ab 24k -ar 44100 -f flv "rtmp://127.0.0.1/live/it"

вещания с локальной веб-камеры
ffmpeg -f video4linux2 -i /dev/video0 -c:v libx264 -an -f flv rtmp://localhost/myapp/mystream

вещаем с веб-камеры с низкой задержкой без звука, рисуем в верхней части картинки текущее время
ffmpeg -r 25 -rtbufsize 1000000k -analyzeduration 0 -s vga -copyts -f dshow -i video="Webcam C170" -vf "drawtext=fontfile=verdana.ttf:fontcolor=yellow@0.8:fontsize=48:box=1:boxcolor=blue@0.8:text=%{localtime}" -s 320x240 -c:v libx264 -g 10 -keyint_min 1 -preset UltraFast -tune zerolatency -crf 25 -an -r 3 -f flv "rtmp://1.2.3.4:1935/live/b.flv live=1"

Wirecast


