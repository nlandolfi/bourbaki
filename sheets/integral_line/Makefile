# generated automatically by bbk sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make integral_line.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

integral_line.tex: sheet.tex macros.tex
	bbk sheet -mode c -in sheet.tex > integral_line.tex

integral_line.pdf: ../../*.tex ../../trademark.pdf *.tex integral_line.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode integral_line.tex
	../../latexrun integral_line.tex
	make terms

graph.csv: sheet.tex
	bbk sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../bbk/static/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: integral_line.pdf
	open integral_line.pdf

o:
	make open

terms: integral_line.tex
	bbk sheet -mode ts integral_line.tex

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
	rm -f integral_line.*
	make reset
	bbk sheet -mode mk > Makefile