import time

def fibonacci(n):
    if n <= 1:
        return n
    return fibonacci(n-1) + fibonacci(n-2)

print("Calculando Fibonacci(45) en Python...")

# Medimos el tiempo de inicio
start_time = time.time()

# Ejecutamos la función
resultado = fibonacci(45)

# Medimos el tiempo de fin
end_time = time.time()

print(f"Resultado: {resultado}")
print(f"Tiempo de ejecución: {end_time - start_time:.4f} segundos")