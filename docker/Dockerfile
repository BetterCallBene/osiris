FROM artifactory.ar.int:5000/ar-python-slim:3.10.10

WORKDIR /usr/local/src

COPY main.py .

# Use the script as the entrypoint
ENTRYPOINT ["python", "main.py"]