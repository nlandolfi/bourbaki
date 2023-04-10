# generated automatically by bbk sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make polynomials.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

polynomials.tex: sheet.tex macros.tex
	bbk sheet -mode c -in sheet.tex > polynomials.tex

polynomials.pdf: ../../*.tex ../../trademark.pdf *.tex polynomials.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode polynomials.tex
	../../latexrun polynomials.tex
	make terms

graph.csv: sheet.tex
	bbk sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../bbk/static/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: polynomials.pdf
	open polynomials.pdf

o:
	make open

terms: polynomials.tex
	bbk sheet -mode ts polynomials.tex

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
	rm -f polynomials.*
	make reset
	bbk sheet -mode mk > Makefile