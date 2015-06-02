attribute vec3 aVertexPosition;
attribute vec2 aUV;
attribute vec3 aNormal;

uniform mat4 uMVMatrix;
uniform mat4 uPMatrix;

varying highp vec2 vUV;
varying highp vec3 vNormal;

void main(void) {
  gl_Position = uPMatrix * uMVMatrix * vec4(aVertexPosition, 1.0);
  vUV = aUV;
	vNormal = vec3(uMVMatrix * vec4(aNormal, 0.0));
}
