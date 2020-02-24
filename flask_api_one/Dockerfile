FROM python:3.7

ENV PYTHONUNBUFFERED=1

WORKDIR /usr/src/flask_api_one

RUN apt-get update && apt-get -y upgrade

COPY Pipfile ./
COPY Pipfile.lock ./

RUN pip3 install pipenv
RUN pipenv install --system
RUN pipenv install --dev
RUN pipenv lock

COPY . ./

CMD ["python", "app.py"]