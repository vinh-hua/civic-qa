FROM python:3.8-slim-buster
ARG PORT=9090
COPY ./requirements.txt /app/
RUN pip install -r /app/requirements.txt
WORKDIR /app
COPY . /app
EXPOSE ${PORT}

CMD ["python", "flask-cog-services/app.py"]