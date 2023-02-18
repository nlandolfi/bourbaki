# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make dynamical_systems.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

dynamical_systems.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > dynamical_systems.tex

dynamical_systems.pdf: ../../*.tex ../../trademark.pdf *.tex dynamical_systems.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode dynamical_systems.tex
	../../latexrun dynamical_systems.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: dynamical_systems.pdf
	open dynamical_systems.pdf

o:
	make open

terms: dynamical_systems.tex
	bbk_sheet -mode ts dynamical_systems.tex

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
	rm -f dynamical_systems.*
	make reset
	bbk_sheet -mode mk > Makefile