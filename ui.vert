attribute vec3 aVertexPosition;
attribute vec2 aUV;

uniform mat4 uMVMatrix;
uniform mat4 uPMatrix;

varying highp vec2 vUV;

void main(void) {
  gl_Position = uPMatrix * uMVMatrix * vec4(aVertexPosition, 1.0);
  vUV = aUV;
}
