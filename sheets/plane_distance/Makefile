# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make plane_distance.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

plane_distance.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > plane_distance.tex

plane_distance.pdf: ../../*.tex ../../trademark.pdf *.tex plane_distance.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode plane_distance.tex
	../../latexrun plane_distance.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: plane_distance.pdf
	open plane_distance.pdf

o:
	make open

terms: plane_distance.tex
	bbk_sheet -mode ts plane_distance.tex

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
	rm -f plane_distance.*
	make reset
	bbk_sheet -mode mk > Makefile