FROM python:3.11-bookworm

WORKDIR /model_service

COPY . .

RUN pip install --upgrade pip
RUN pip install -r requirements.txt

CMD ["python", "model_service.py"]