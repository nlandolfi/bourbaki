# generated automatically by bbk sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make orthogonal_complements.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

orthogonal_complements.tex: sheet.tex macros.tex
	bbk sheet -mode c -in sheet.tex > orthogonal_complements.tex

orthogonal_complements.pdf: ../../*.tex ../../trademark.pdf *.tex orthogonal_complements.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode orthogonal_complements.tex
	../../latexrun orthogonal_complements.tex
	make terms

graph.csv: sheet.tex
	bbk sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../bbk/static/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: orthogonal_complements.pdf
	open orthogonal_complements.pdf

o:
	make open

terms: orthogonal_complements.tex
	bbk sheet -mode ts orthogonal_complements.tex

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
	rm -f orthogonal_complements.*
	make reset
	bbk sheet -mode mk > Makefile