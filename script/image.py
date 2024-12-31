from PIL import Image

def extract_watermark(image_path):
    img = Image.open(image_path)
    pixels = img.load()
    watermark = ""
    for i in range(img.width):
        for j in range(img.height):
            r, g, b = pixels[i, j]
            watermark += str(r & 1)
            watermark += str(g & 1)
            watermark += str(b & 1)
    return watermark

image_path = "your_image_with_watermark.jpg"
extracted_watermark = extract_watermark(image_path)
print(extracted_watermark)