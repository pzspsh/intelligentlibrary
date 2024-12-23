from django.shortcuts import render, HttpResponse

# Create your views here.


def index(request):
    return HttpResponse("欢迎使用Django")