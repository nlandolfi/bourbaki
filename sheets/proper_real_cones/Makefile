# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make proper_real_cones.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

proper_real_cones.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > proper_real_cones.tex

proper_real_cones.pdf: ../../*.tex ../../trademark.pdf *.tex proper_real_cones.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode proper_real_cones.tex
	../../latexrun proper_real_cones.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: proper_real_cones.pdf
	open proper_real_cones.pdf

o:
	make open

terms: proper_real_cones.tex
	bbk_sheet -mode ts proper_real_cones.tex

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
	rm -f proper_real_cones.*
	make reset
	bbk_sheet -mode mk > Makefile