if [[ "$OSTYPE" == "linux-gnu"* ]]; then
    if [ ! -d "venv" ]; then
        python3 -m venv venv
    fi
    source venv/bin/activate
elif [[ "$OSTYPE" == "msys"* ]]; then
    if [ ! -d "venv" ]; then
        python -m venv venv
    fi
    venv\Scripts\activate
fi

pip install setuptools
pip install -r requirements.txt

if [ ! -d "models" ]; then
    mkdir models
fi

if [ ! -d "dataset" ]; then
    mkdir dataset
fi

wget https://drive.google.com/uc?export=download&id=1oiENbYlhfiq2J6zHUiULx9XptdOV3-UY -O models/doodle-guessr.ph
wget https://drive.google.com/uc?export=download&id=1ScYCfoo7Khw2aTUmNNVY5c9qWRF7K82N -O dataset/doodle-dataset.json

