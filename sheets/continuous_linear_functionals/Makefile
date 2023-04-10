# generated automatically by bbk sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make continuous_linear_functionals.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

continuous_linear_functionals.tex: sheet.tex macros.tex
	bbk sheet -mode c -in sheet.tex > continuous_linear_functionals.tex

continuous_linear_functionals.pdf: ../../*.tex ../../trademark.pdf *.tex continuous_linear_functionals.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode continuous_linear_functionals.tex
	../../latexrun continuous_linear_functionals.tex
	make terms

graph.csv: sheet.tex
	bbk sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../bbk/static/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: continuous_linear_functionals.pdf
	open continuous_linear_functionals.pdf

o:
	make open

terms: continuous_linear_functionals.tex
	bbk sheet -mode ts continuous_linear_functionals.tex

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
	rm -f continuous_linear_functionals.*
	make reset
	bbk sheet -mode mk > Makefile