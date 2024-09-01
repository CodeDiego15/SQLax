# SQLax

SQLax es una herramienta para la detección y explotación de vulnerabilidades de inyección SQL en aplicaciones web. 

## Características

- **Detección**: Identifica posibles vulnerabilidades de inyección SQL en una URL.
- **Explotación**: Intenta explotar las vulnerabilidades detectadas para verificar su existencia.

## Instalación

Para utilizar SQLax, necesitas tener [Go](https://golang.org/dl/) instalado en tu sistema. Sigue estos pasos para instalar y usar SQLax:

1. **Clona el repositorio Macos y Linux**:
   ```bash
   git clone https://github.com/CodeDiego15/SQLax.git
   cd SQLax
   go build -o sqlax
   sudo mv sqlax /usr/local/bin   
   ```

## Uso 
El comando principal para ejecutar SQLax es:
```bash
sqlax -u <URL> -v -e
```
1. Detectar vulnerabilidades SQL:
```bash
sqlax -u http://example.com/page.php?id=1 -v
```
2. Explotar vulnerabilidades SQL:
```bash
sqlax -u http://example.com/page.php?id=1 -e
```
3. Detección y explotación en un solo comando:
```bash
sqlax -u http://example.com/page.php?id=1 -v -e
```
