# generated automatically by bbk sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make ordinary_row_reductions.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

ordinary_row_reductions.tex: sheet.tex macros.tex
	bbk sheet -mode c -in sheet.tex > ordinary_row_reductions.tex

ordinary_row_reductions.pdf: ../../*.tex ../../trademark.pdf *.tex ordinary_row_reductions.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode ordinary_row_reductions.tex
	../../latexrun ordinary_row_reductions.tex
	make terms

graph.csv: sheet.tex
	bbk sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../bbk/static/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: ordinary_row_reductions.pdf
	open ordinary_row_reductions.pdf

o:
	make open

terms: ordinary_row_reductions.tex
	bbk sheet -mode ts ordinary_row_reductions.tex

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
	rm -f ordinary_row_reductions.*
	make reset
	bbk sheet -mode mk > Makefile