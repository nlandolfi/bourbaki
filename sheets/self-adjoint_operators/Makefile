# generated automatically by bbk sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make self-adjoint_operators.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

self-adjoint_operators.tex: sheet.tex macros.tex
	bbk sheet -mode c -in sheet.tex > self-adjoint_operators.tex

self-adjoint_operators.pdf: ../../*.tex ../../trademark.pdf *.tex self-adjoint_operators.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode self-adjoint_operators.tex
	../../latexrun self-adjoint_operators.tex
	make terms

graph.csv: sheet.tex
	bbk sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../bbk/static/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: self-adjoint_operators.pdf
	open self-adjoint_operators.pdf

o:
	make open

terms: self-adjoint_operators.tex
	bbk sheet -mode ts self-adjoint_operators.tex

spell:
	aspell -c sheet.tex

s:
	make spell

reset:
	rm -f graph.pdf
	rm -f graph.csv
	rm -f graph.graphviz

clean:
	make reset

remake:
	rm -f self-adjoint_operators.*
	make reset
	bbk sheet -mode mk > Makefile