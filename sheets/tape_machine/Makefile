# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make tape_machine.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

tape_machine.tex: sheet.tex macros.tex
	bbk_sheet -mode c -in sheet.tex > tape_machine.tex

tape_machine.pdf: ../../*.tex ../../trademark.pdf *.tex tape_machine.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode tape_machine.tex
	../../latexrun tape_machine.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: tape_machine.pdf
	open tape_machine.pdf

o:
	make open

terms: tape_machine.tex
	bbk_sheet -mode ts tape_machine.tex

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
	rm -f tape_machine.*
	make reset
	bbk_sheet -mode mk > Makefile