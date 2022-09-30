# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make cubes.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

cubes.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > cubes.tex

cubes.pdf: ../../*.tex ../../trademark.pdf *.tex cubes.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode cubes.tex
	../../latexrun cubes.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: cubes.pdf
	open cubes.pdf

o:
	make open

terms: cubes.tex
	bbk_sheet -mode ts cubes.tex

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
	rm -f cubes.*
	make reset
	bbk_sheet -mode mk > Makefile