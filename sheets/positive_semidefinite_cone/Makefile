# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make positive_semidefinite_cone.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

positive_semidefinite_cone.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > positive_semidefinite_cone.tex

positive_semidefinite_cone.pdf: ../../*.tex ../../trademark.pdf *.tex positive_semidefinite_cone.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode positive_semidefinite_cone.tex
	../../latexrun positive_semidefinite_cone.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: positive_semidefinite_cone.pdf
	open positive_semidefinite_cone.pdf

o:
	make open

terms: positive_semidefinite_cone.tex
	bbk_sheet -mode ts positive_semidefinite_cone.tex

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
	rm -f positive_semidefinite_cone.*
	make reset
	bbk_sheet -mode mk > Makefile