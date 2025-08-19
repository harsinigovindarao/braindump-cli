import grpc
from concurrent import futures
import time

import thoughts_pb2
import thoughts_pb2_grpc

from transformers import AutoTokenizer, AutoModelForSequenceClassification
import torch
import numpy as np

# Load sentiment model (Tone)
MODEL = "cardiffnlp/twitter-roberta-base-sentiment"
tokenizer = AutoTokenizer.from_pretrained(MODEL)
model = AutoModelForSequenceClassification.from_pretrained(MODEL)

# Labels
labels = ['Negative', 'Neutral', 'Positive']

# Simple category keywords for demo (you can improve this)
CATEGORY_MAP = {
    "health": ["doctor", "hospital", "medicine", "pain"],
    "work": ["meeting", "deadline", "office", "project"],
    "finance": ["money", "loan", "budget", "savings"],
    "learning": ["course", "exam", "study", "book"],
    "random": []
}


def classify_category(text):
    for category, keywords in CATEGORY_MAP.items():
        for kw in keywords:
            if kw.lower() in text.lower():
                return category
    return "random"


def classify_tone(text):
    inputs = tokenizer(text, return_tensors="pt", truncation=True)
    with torch.no_grad():
        logits = model(**inputs).logits
    scores = torch.nn.functional.softmax(logits, dim=1)
    predicted = torch.argmax(scores, dim=1).item()
    return labels[predicted]


class ThoughtServiceServicer(thoughts_pb2_grpc.ThoughtServiceServicer):
    def ClassifyThought(self, request, context):
        text = request.text
        category = classify_category(text)
        tone = classify_tone(text)
        return thoughts_pb2.ThoughtResponse(category=category, tone=tone)


def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=5))
    thoughts_pb2_grpc.add_ThoughtServiceServicer_to_server(ThoughtServiceServicer(), server)
    server.add_insecure_port('[::]:50051')
    print("🧠 Python NLP gRPC server running on port 50051...")
    server.start()
    try:
        while True:
            time.sleep(86400)
    except KeyboardInterrupt:
        print("🔌 Shutting down gRPC server")
        server.stop(0)


if __name__ == "__main__":
    serve()
